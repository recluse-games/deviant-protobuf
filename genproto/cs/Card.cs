// <auto-generated>
//     Generated by the protocol buffer compiler.  DO NOT EDIT!
//     source: Card.proto
// </auto-generated>
#pragma warning disable 1591, 0612, 3021
#region Designer generated code

using pb = global::Google.Protobuf;
using pbc = global::Google.Protobuf.Collections;
using pbr = global::Google.Protobuf.Reflection;
using scg = global::System.Collections.Generic;
namespace Deviant {

  /// <summary>Holder for reflection information generated from Card.proto</summary>
  public static partial class CardReflection {

    #region Descriptor
    /// <summary>File descriptor for Card.proto</summary>
    public static pbr::FileDescriptor Descriptor {
      get { return descriptor; }
    }
    private static pbr::FileDescriptor descriptor;

    static CardReflection() {
      byte[] descriptorData = global::System.Convert.FromBase64String(
          string.Concat(
            "CgpDYXJkLnByb3RvEgdEZXZpYW50GhBDYXJkQWN0aW9uLnByb3RvGg5DYXJk",
            "VHlwZS5wcm90byKaAQoEQ2FyZBIKCgJpZBgBIAEoCRIMCgRjb3N0GAIgASgF",
            "Eg0KBXRpdGxlGAMgASgJEg4KBmZsYXZvchgEIAEoCRITCgtkZXNjcmlwdGlv",
            "bhgFIAEoCRIfCgR0eXBlGAYgASgOMhEuRGV2aWFudC5DYXJkVHlwZRIjCgZh",
            "Y3Rpb24YByABKAsyEy5EZXZpYW50LkNhcmRBY3Rpb25CCVoHZGV2aWFudGIG",
            "cHJvdG8z"));
      descriptor = pbr::FileDescriptor.FromGeneratedCode(descriptorData,
          new pbr::FileDescriptor[] { global::Deviant.CardActionReflection.Descriptor, global::Deviant.CardTypeReflection.Descriptor, },
          new pbr::GeneratedClrTypeInfo(null, new pbr::GeneratedClrTypeInfo[] {
            new pbr::GeneratedClrTypeInfo(typeof(global::Deviant.Card), global::Deviant.Card.Parser, new[]{ "Id", "Cost", "Title", "Flavor", "Description", "Type", "Action" }, null, null, null)
          }));
    }
    #endregion

  }
  #region Messages
  public sealed partial class Card : pb::IMessage<Card> {
    private static readonly pb::MessageParser<Card> _parser = new pb::MessageParser<Card>(() => new Card());
    private pb::UnknownFieldSet _unknownFields;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pb::MessageParser<Card> Parser { get { return _parser; } }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pbr::MessageDescriptor Descriptor {
      get { return global::Deviant.CardReflection.Descriptor.MessageTypes[0]; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    pbr::MessageDescriptor pb::IMessage.Descriptor {
      get { return Descriptor; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public Card() {
      OnConstruction();
    }

    partial void OnConstruction();

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public Card(Card other) : this() {
      id_ = other.id_;
      cost_ = other.cost_;
      title_ = other.title_;
      flavor_ = other.flavor_;
      description_ = other.description_;
      type_ = other.type_;
      action_ = other.action_ != null ? other.action_.Clone() : null;
      _unknownFields = pb::UnknownFieldSet.Clone(other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public Card Clone() {
      return new Card(this);
    }

    /// <summary>Field number for the "id" field.</summary>
    public const int IdFieldNumber = 1;
    private string id_ = "";
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public string Id {
      get { return id_; }
      set {
        id_ = pb::ProtoPreconditions.CheckNotNull(value, "value");
      }
    }

    /// <summary>Field number for the "cost" field.</summary>
    public const int CostFieldNumber = 2;
    private int cost_;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public int Cost {
      get { return cost_; }
      set {
        cost_ = value;
      }
    }

    /// <summary>Field number for the "title" field.</summary>
    public const int TitleFieldNumber = 3;
    private string title_ = "";
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public string Title {
      get { return title_; }
      set {
        title_ = pb::ProtoPreconditions.CheckNotNull(value, "value");
      }
    }

    /// <summary>Field number for the "flavor" field.</summary>
    public const int FlavorFieldNumber = 4;
    private string flavor_ = "";
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public string Flavor {
      get { return flavor_; }
      set {
        flavor_ = pb::ProtoPreconditions.CheckNotNull(value, "value");
      }
    }

    /// <summary>Field number for the "description" field.</summary>
    public const int DescriptionFieldNumber = 5;
    private string description_ = "";
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public string Description {
      get { return description_; }
      set {
        description_ = pb::ProtoPreconditions.CheckNotNull(value, "value");
      }
    }

    /// <summary>Field number for the "type" field.</summary>
    public const int TypeFieldNumber = 6;
    private global::Deviant.CardType type_ = 0;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public global::Deviant.CardType Type {
      get { return type_; }
      set {
        type_ = value;
      }
    }

    /// <summary>Field number for the "action" field.</summary>
    public const int ActionFieldNumber = 7;
    private global::Deviant.CardAction action_;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public global::Deviant.CardAction Action {
      get { return action_; }
      set {
        action_ = value;
      }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override bool Equals(object other) {
      return Equals(other as Card);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public bool Equals(Card other) {
      if (ReferenceEquals(other, null)) {
        return false;
      }
      if (ReferenceEquals(other, this)) {
        return true;
      }
      if (Id != other.Id) return false;
      if (Cost != other.Cost) return false;
      if (Title != other.Title) return false;
      if (Flavor != other.Flavor) return false;
      if (Description != other.Description) return false;
      if (Type != other.Type) return false;
      if (!object.Equals(Action, other.Action)) return false;
      return Equals(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override int GetHashCode() {
      int hash = 1;
      if (Id.Length != 0) hash ^= Id.GetHashCode();
      if (Cost != 0) hash ^= Cost.GetHashCode();
      if (Title.Length != 0) hash ^= Title.GetHashCode();
      if (Flavor.Length != 0) hash ^= Flavor.GetHashCode();
      if (Description.Length != 0) hash ^= Description.GetHashCode();
      if (Type != 0) hash ^= Type.GetHashCode();
      if (action_ != null) hash ^= Action.GetHashCode();
      if (_unknownFields != null) {
        hash ^= _unknownFields.GetHashCode();
      }
      return hash;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override string ToString() {
      return pb::JsonFormatter.ToDiagnosticString(this);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void WriteTo(pb::CodedOutputStream output) {
      if (Id.Length != 0) {
        output.WriteRawTag(10);
        output.WriteString(Id);
      }
      if (Cost != 0) {
        output.WriteRawTag(16);
        output.WriteInt32(Cost);
      }
      if (Title.Length != 0) {
        output.WriteRawTag(26);
        output.WriteString(Title);
      }
      if (Flavor.Length != 0) {
        output.WriteRawTag(34);
        output.WriteString(Flavor);
      }
      if (Description.Length != 0) {
        output.WriteRawTag(42);
        output.WriteString(Description);
      }
      if (Type != 0) {
        output.WriteRawTag(48);
        output.WriteEnum((int) Type);
      }
      if (action_ != null) {
        output.WriteRawTag(58);
        output.WriteMessage(Action);
      }
      if (_unknownFields != null) {
        _unknownFields.WriteTo(output);
      }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public int CalculateSize() {
      int size = 0;
      if (Id.Length != 0) {
        size += 1 + pb::CodedOutputStream.ComputeStringSize(Id);
      }
      if (Cost != 0) {
        size += 1 + pb::CodedOutputStream.ComputeInt32Size(Cost);
      }
      if (Title.Length != 0) {
        size += 1 + pb::CodedOutputStream.ComputeStringSize(Title);
      }
      if (Flavor.Length != 0) {
        size += 1 + pb::CodedOutputStream.ComputeStringSize(Flavor);
      }
      if (Description.Length != 0) {
        size += 1 + pb::CodedOutputStream.ComputeStringSize(Description);
      }
      if (Type != 0) {
        size += 1 + pb::CodedOutputStream.ComputeEnumSize((int) Type);
      }
      if (action_ != null) {
        size += 1 + pb::CodedOutputStream.ComputeMessageSize(Action);
      }
      if (_unknownFields != null) {
        size += _unknownFields.CalculateSize();
      }
      return size;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void MergeFrom(Card other) {
      if (other == null) {
        return;
      }
      if (other.Id.Length != 0) {
        Id = other.Id;
      }
      if (other.Cost != 0) {
        Cost = other.Cost;
      }
      if (other.Title.Length != 0) {
        Title = other.Title;
      }
      if (other.Flavor.Length != 0) {
        Flavor = other.Flavor;
      }
      if (other.Description.Length != 0) {
        Description = other.Description;
      }
      if (other.Type != 0) {
        Type = other.Type;
      }
      if (other.action_ != null) {
        if (action_ == null) {
          Action = new global::Deviant.CardAction();
        }
        Action.MergeFrom(other.Action);
      }
      _unknownFields = pb::UnknownFieldSet.MergeFrom(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void MergeFrom(pb::CodedInputStream input) {
      uint tag;
      while ((tag = input.ReadTag()) != 0) {
        switch(tag) {
          default:
            _unknownFields = pb::UnknownFieldSet.MergeFieldFrom(_unknownFields, input);
            break;
          case 10: {
            Id = input.ReadString();
            break;
          }
          case 16: {
            Cost = input.ReadInt32();
            break;
          }
          case 26: {
            Title = input.ReadString();
            break;
          }
          case 34: {
            Flavor = input.ReadString();
            break;
          }
          case 42: {
            Description = input.ReadString();
            break;
          }
          case 48: {
            Type = (global::Deviant.CardType) input.ReadEnum();
            break;
          }
          case 58: {
            if (action_ == null) {
              Action = new global::Deviant.CardAction();
            }
            input.ReadMessage(Action);
            break;
          }
        }
      }
    }

  }

  #endregion

}

#endregion Designer generated code